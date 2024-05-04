import { Component } from '@angular/core';
import { UserService } from '../services/user.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-user',
  //standalone: true,
  //imports: [],
  templateUrl: './user.component.html',
  styleUrl: './user.component.css'
})
export class UserComponent {
  user: any;

  constructor(private authService: AuthService, private userService: UserService) {}

  ngOnInit(): void {
    const token = this.authService.getToken();
    if (token) {
      this.getUserProfile();
    }
  }

  getUserProfile(): void {
    this.userService.getUserProfile()
      .subscribe(
        (data) => {
          this.user = data;
        },
        (error) => {
          console.error('Error fetching user profile:', error);
        }
      );
  }
}

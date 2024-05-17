import { Component, OnInit } from '@angular/core';
import { UserService } from '../services/user.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  user: any;

  constructor(private userService: UserService, private authService: AuthService) {}

  ngOnInit(): void {
    const token = this.authService.getToken();
    if (!token) {
      // Если токен отсутствует, перенаправляем пользователя на страницу входа
      // Реализация редиректа зависит от вашего роутинга
      return;
    }
    this.getUserProfile();
  }

  getUserProfile(): void {
    this.userService.getUserProfile().subscribe(
      (data) => {
        console.log('User profile data:', data);
        if (data) {
          this.user = data;
        } else {
          console.error('User profile data is empty or undefined.');
        }
      },
      (error) => {
        console.error('Error fetching user profile:', error);
      }
    );
  }
}

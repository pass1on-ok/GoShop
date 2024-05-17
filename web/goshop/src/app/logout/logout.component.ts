// logout.component.ts

import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.css']
})
export class LogoutComponent {
/*
  constructor(private authService: AuthService, private router: Router) {}

  logoutUser() {
    
    this.authService.logout()
      .subscribe(
        () => {
          
          localStorage.removeItem('token');
          // Перенаправляем пользователя на страницу входа
          this.router.navigate(['/login']);
        },
        (error) => {
          console.error('An error occurred while logging out:', error);
        }
      );
  }
  */
}

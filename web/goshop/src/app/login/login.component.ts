// login.component.ts

import { Component } from '@angular/core';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  constructor(private authService: AuthService) { }

  login(email: string, password: string): void {
    this.authService.login(email, password).subscribe(
      
      res => console.log('Logged in successfully'),
      
      err => console.error('Login error:', err)
    );
  }
}

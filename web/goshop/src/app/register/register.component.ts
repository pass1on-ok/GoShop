import { Component } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {
  username!: string;
  email: string = '';
  password: string = '';
  confirmPassword!: string;
  termsAccepted: boolean = false;

  constructor(private authService: AuthService,
    private router: Router
  ) {}

  registerUser() {
    if (!this.validateUsername(this.username)) {
      alert('Please enter a valid username.');
      return;
    }
    if (!this.validateEmail(this.email)) {
      alert('Please enter a valid email address.');
      return;
    }
    if (!this.validatePassword(this.password)) {
      alert('Password must be at least 6 characters long.');
      return;
    }
    if (this.password !== this.confirmPassword) {
      alert('Passwords do not match.');
      return;
    }
    if (!this.termsAccepted) {
      alert('Please accept the terms & conditions.');
      return;
    }

    
    this.authService.register(this.username, this.email, this.password)
      .subscribe(
        response => {
          
          console.log('Registration successful:', response);
          this.router.navigate(['/login']);

        },
        error => {
          
          console.error('Registration failed:', error);
          alert('Registration failed. Please try again.');
        }
      );
  }

  validateUsername(username: string): boolean {
    return username.length > 0;
  }

  validateEmail(email: string): boolean {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  }

  validatePassword(password: string): boolean {

    return password.length >= 6;
  }
}

import { Component, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.css'],
  providers: [AuthService],
})
export class ToolbarComponent {
  title = 'toolbar';
  constructor(
    private auth: AuthService,
    private router: Router,
  ) {}
  home(): void {
    this.router.navigateByUrl('/home');
  }
  login(): void {
    this.router.navigateByUrl('/login');
  }
  logout(): void {
    this.router.navigateByUrl('/logout');
  }
  signUp(): void {
    this.router.navigateByUrl('/register');
  }
  newEntry(): void {
    this.router.navigateByUrl('/newentry');
  }
}

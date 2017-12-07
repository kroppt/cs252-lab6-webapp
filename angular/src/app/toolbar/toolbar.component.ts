import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.css']
})
export class ToolbarComponent implements  OnInit {
  title = 'toolbar';
  isLoggedIn: boolean = null;
  constructor( private cookieService: CookieService, private router: Router ) {}
  ngOnInit(): void {
    if (this.cookieService.get('Name') === 'Auth' &&
      this.cookieService.get('Value') != null) {
        this.isLoggedIn = true;
    } else {
      this.isLoggedIn = false;
    }
  }
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

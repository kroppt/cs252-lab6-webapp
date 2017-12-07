import { Component } from '@angular/core';
import { Http } from '@angular/http';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
})
export class AccountComponent {
  title = 'account';
  constructor(private http: Http) { }
}

import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

@Injectable()
export class ServerService {
  apiRoot = '35.227.48.112/api';
  constructor(private http: Http) { }
  authUser(username, password): string {
    this.http
      .post(this.apiRoot + '/authUser',
      { username: username, password: btoa(password) })
      .subscribe();
    return '';
  }
  loginUser(username, password): string {
    this.http
      .post(this.apiRoot + '/loginUser', { username: username })
      .subscribe();
    return '';
  }
  logoutUser(): string {
    this.http
      .post(this.apiRoot + '/logoutUser', {})
      .subscribe();
    return '';
  }
  newUser(username, password): string {
    this.http
      .post(this.apiRoot + '/newUser',
      { username: username, password: btoa(password) })
      .subscribe();
    return '';
  }
}

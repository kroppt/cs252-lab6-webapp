import { Injectable } from '@angular/core';

@Injectable()
export class AuthService {
    constructor() { }
    getUsername() {
        return localStorage.getItem('currentUsername');
    }
    isLoggedIn() {
        return localStorage.hasOwnProperty('currentUsername');
    }
    logIn(username: string) {
        localStorage.setItem('currentUsername', username);
    }
    logOut() {
        localStorage.removeItem('currentUsername');
    }
}

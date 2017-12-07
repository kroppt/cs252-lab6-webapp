import { Component } from '@angular/core';
import {
  FormGroup,
  FormControl,
  Validators,
} from '@angular/forms';
import { Http } from '@angular/http';
import { passwordMatchValidator } from '../shared/password-match.directive';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
})
export class RegisterComponent {
  title = 'register';

  constructor(private http: Http) { }

  username = new FormControl(
    '',
    [
      Validators.required,
      Validators.minLength(6),
      Validators.maxLength(20),
    ]
   );

   password = new FormControl(
    '',
    [
      Validators.required,
      Validators.minLength(10),
      Validators.maxLength(50),
    ]
  );

  passwordVerify = new FormControl(
    '',
    [
      passwordMatchValidator(this.password),
    ]
  );

  submit() {
    this.http
      .post('35.227.48.112/api/newUser',
      {
        username: this.username.value,
        password: btoa(this.password.value),
      })
      .subscribe();
    return;
  }

}

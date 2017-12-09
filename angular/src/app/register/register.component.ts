import { Component } from '@angular/core';
import {
  FormsModule,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { Http } from '@angular/http';
import { passwordMatchValidator } from '../shared/password-match.directive';
import { MatSnackBar } from '@angular/material';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
})
export class RegisterComponent {
  title = 'register';

  constructor(private http: Http, private snackBar: MatSnackBar) { }

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
      .post('http://35.227.48.112/api/newUser',
      {
        username: this.username.value,
        password: btoa(this.password.value),
      })
      .subscribe(
        data => {
          console.log(data);
        },
        err => {
          this.snackBar.open(err._body, 'OK', {
            duration: 2000,
          });
        });
    console.log();
    return;
  }

}

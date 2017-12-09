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
import { MatSnackBar } from '@angular/material';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent {
  title = 'login';

  constructor(private router:Router, private http: Http, private snackBar: MatSnackBar) { }

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

  submit() {
    this.http
      .post('http://35.227.48.112/api/loginUser',
      {
        username: this.username.value,
        password: btoa(this.password.value),
      })
      .subscribe(
      data => {
        this.router.navigate(['/']);
      },
      err => {
        let message: string;
        if (err.status === 0) {
          message = 'Connection error';
        } else if (err.status / 500 >= 1) {
          message = 'Server error';
        } else {
          message = err._body;
        }
        this.snackBar.open(message, 'OK', {
          duration: 2000,
        });
      });
    return;
  }
}

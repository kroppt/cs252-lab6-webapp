import { Component, OnInit } from '@angular/core';
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
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  title = 'login';

  constructor(
    private auth: AuthService,
    private router: Router,
    private http: Http,
    private snackBar: MatSnackBar,
  ) { }

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

  isFormInvalid() {
    return this.username.invalid || this.password.invalid;
  }

  submit() {
    if (this.isFormInvalid()) {
      return;
    }
    this.http.post('http://35.227.48.112/api/loginUser',
      {
        username: this.username.value,
        password: btoa(this.password.value),
      },
      {
        withCredentials: true
      },
    ).subscribe(
      data => {
        this.auth.logIn(this.username.value);
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

  ngOnInit() {
    if (this.auth.isLoggedIn()) {
      this.router.navigate(['/']);
    }
  }

}

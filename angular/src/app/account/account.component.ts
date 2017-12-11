import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth/auth.service';
import { Http } from '@angular/http';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material';
import { MatExpansionModule, MatAccordion } from '@angular/material/expansion';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { passwordMatchValidator } from '../shared/password-match.directive';
import {
  FormsModule,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css'],
})
export class AccountComponent implements OnInit {
  title = 'account';

  constructor(
    private auth: AuthService,
    private router: Router,
    private http: Http,
    private snackBar: MatSnackBar,
  ) { }

  username = new FormControl(
    {value: this.auth.getUsername(), disabled: true},
    []
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

  isFormInvalid() {
    return this.password.invalid || this.passwordVerify.invalid;
  }

  submit() {
    if (this.isFormInvalid()) {
      return;
    }
    this.http.post('http://35.227.48.112/api/changePassword',
      {
        username: '',
        password: btoa(this.password.value),
      },
      {
        withCredentials: true
      },
    ).subscribe(
      data => {
        this.snackBar.open('Password successfully changed', 'OK', {
          duration: 2000,
        });
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
    if (!this.auth.isLoggedIn()) {
      this.router.navigate(['/']);
    }
  }

}

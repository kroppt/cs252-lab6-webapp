import { Component } from '@angular/core';
import {
  FormControl,
  Validators,
} from '@angular/forms';
import { Http } from '@angular/http';

@Component({
  selector: 'app-new-entry',
  templateUrl: './new-entry.component.html',
})
export class NewEntryComponent {
  title = 'new-entry';

  constructor(private http: Http) { }

  name = new FormControl(
    '',
    [
      Validators.required,
      Validators.minLength(6),
    ]
   );

   url = new FormControl(
    '',
    [
      Validators.required,
    ]
  );

  submit() {
    this.http
      .post('35.227.48.112/api/newEntry',
      {
        name: this.name.value,
        url: btoa(this.url.value),
      })
      .subscribe();
    return;
  }
}

import { Component, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  encapsulation: ViewEncapsulation.None,
  preserveWhitespaces: false,
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';
}

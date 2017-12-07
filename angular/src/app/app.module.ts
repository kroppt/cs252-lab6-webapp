import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { SignupComponent } from './signup/signup.component';
import { HttpModule } from '@angular/http';
import 'hammerjs';

import {
  MatButtonModule,
  MatIconModule,
  MatToolbarModule
} from '@angular/material';

@NgModule({
  exports: [
    MatButtonModule,
    MatIconModule,
    MatToolbarModule
  ]
})
export class MaterialModule {}

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    SignupComponent
  ],
  imports: [
    BrowserModule,
    HttpModule,
    MaterialModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}

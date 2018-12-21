import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CaptchaComponent } from './captcha/captcha.component';

@NgModule({
  declarations: [
    AppComponent,
    CaptchaComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    HttpClientModule

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

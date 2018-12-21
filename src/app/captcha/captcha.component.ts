import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Captcha } from '../captcha';

@Component({
  selector: 'app-captcha',
  templateUrl: './captcha.component.html',
  styleUrls: ['./captcha.component.css']
})
export class CaptchaComponent implements OnInit {

  public capt:Captcha;

  constructor(
    private http: HttpClient
  ) { }

  ngOnInit() {
    this.getnaja();
  }

  public getnaja(): void {
    this.http.get('http://localhost:1323/getCaptcha')
      .subscribe(data => {
        this.capt=data['captcha']
      })
  }

}
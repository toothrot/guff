import {Component, OnInit} from '@angular/core';
import {AppService} from "./app.service";
import {Observable} from "rxjs";
import {GetCurrentUserResponse} from 'src/generated/users_pb';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less'],
})
export class AppComponent implements OnInit {
  title = 'web';
  currentUser: Observable<GetCurrentUserResponse>;

  constructor(private appService: AppService) {
  }

  ngOnInit() {
    this.currentUser = this.appService.getCurrentUser();
  }
}

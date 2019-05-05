import {Component, OnInit} from '@angular/core';
import {AppService} from './app.service';
import {Observable, from, of} from 'rxjs';
import {GetCurrentUserResponse} from 'src/generated/users_pb';
import {ActivatedRoute} from '@angular/router';
import {map, mergeMap, tap, catchError} from 'rxjs/operators';
import {OAuthService} from './oauth/oauth.service';
import {filter} from 'rxjs/internal/operators/filter';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less'],
})
export class AppComponent implements OnInit {
  title = 'web';
  currentUser: Observable<GetCurrentUserResponse>;
  loginURL: Observable<string>;

  constructor(private appService: AppService, private route: ActivatedRoute, private oauthService: OAuthService) {
  }

  ngOnInit() {
    this.currentUser = this.route.fragment.pipe(
      map(fragment => new URLSearchParams(fragment)),
      // mergeMap to request to validate token on backend.
      // if successful, save token to local storage and redirect
      // on failure, redirect and flash error.
      map(params => params.get('id_token')),
      filter(token => {
        console.log(token);
        return !!token;
      }),
      mergeMap(token => this.oauthService.verifyAndSave(token)),
      // map(user => user.getEmail()),
      // catchError(err => of(err.toString())),
      // tap(() => from(this.router.navigateByUrl('/'))),
    );

    // this.currentUser = this.appService.getCurrentUser();
    this.loginURL = this.appService.oAuthURL();
  }
}

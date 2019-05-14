import {Component, OnInit} from '@angular/core';
import {AppService} from './app.service';
import {from, Observable} from 'rxjs';
import {GetCurrentUserResponse} from 'src/generated/users_pb';
import {ActivatedRoute, Router} from '@angular/router';
import {concatMap, map, mergeMap, shareReplay, tap} from 'rxjs/operators';
import {OAuthService} from './oauth/oauth.service';
import {filter} from 'rxjs/internal/operators/filter';
import {GetDivisionsResponse} from '../generated/divisions_pb';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less'],
})
export class AppComponent implements OnInit {
  title = 'web';
  currentUser: Observable<GetCurrentUserResponse>;
  divisions: Observable<GetDivisionsResponse>;
  loginURL: Observable<string>;

  constructor(private appService: AppService, private route: ActivatedRoute, private oauthService: OAuthService, private router: Router) {
    this.loginURL = this.appService.oAuthURL().pipe(shareReplay(1));
    this.divisions = this.appService.getDivisions().pipe(shareReplay(1));
  }

  ngOnInit() {
    this.currentUser = this.route.fragment.pipe(
      map(fragment => new URLSearchParams(fragment)),
      // mergeMap to request to validate token on backend.
      // if successful, save token to local storage and redirect
      // on failure, redirect and flash error.
      map(params => params.get('id_token') || localStorage.getItem('token')),
      filter(token => {
        console.log(token);
        return !!token;
      }),
      mergeMap(token => this.oauthService.verifyAndSave(token)),
      // map(user => user.getEmail()),
      // catchError(err => of(err.toString())),
      tap(() => from(this.router.navigateByUrl('/'))),
    );
  }

  scrape() {
    this.divisions = this.appService.scrapeDivisions().pipe(
      concatMap(() => {
        return this.appService.getDivisions().pipe(shareReplay(1));
      })
    );
  }
}

import {Component, OnInit} from '@angular/core';
import {from, Observable, of} from 'rxjs';
import {ActivatedRoute, Router} from '@angular/router';
import {catchError, map, mergeMap, tap} from 'rxjs/operators';
import {OAuthCallbackService} from './oauth-callback.service';

@Component({
  selector: 'app-oauth-callback',
  templateUrl: './oauth-callback.component.html',
  styleUrls: ['./oauth-callback.component.less'],
})
export class OAuthCallbackComponent implements OnInit {
  fragment = new Observable<string>();

  constructor(private route: ActivatedRoute, private service: OAuthCallbackService, private router: Router) {
  }

  ngOnInit() {
    this.fragment = this.route.fragment.pipe(
      map(fragment => new URLSearchParams(fragment)),
      // mergeMap to request to validate token on backend.
      // if successful, save token to local storage and redirect
      // on failure, redirect and flash error.
      map(params => params.get('id_token')),
      mergeMap(token => this.service.verifyAndSave(token)),
      map(user => user.getEmail()),
      catchError(err => of(err.toString())),
      tap(() => from(this.router.navigateByUrl('/'))),
    );
  }

}

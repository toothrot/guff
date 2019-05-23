import { Injectable } from '@angular/core';
import {UsersServiceClient} from '../../generated/UsersServiceClientPb';
import {AdminServiceClient} from '../../generated/AdminServiceClientPb';
import {bindNodeCallback, Observable} from 'rxjs';
import {GetCurrentUserRequest, GetCurrentUserResponse} from '../../generated/users_pb';
import {AppService} from '../app.service';
import {map, shareReplay, take} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class UsersService {
  client: UsersServiceClient;
  adminClient: AdminServiceClient;

  constructor() {
    this.client = new UsersServiceClient('', {}, {});
    this.adminClient = new AdminServiceClient('', {}, {});
  }

  private static authHeader() {
    return {Authorization: 'Bearer ' + localStorage.getItem('token')};
  }

  getCurrentUser(): Observable<GetCurrentUserResponse> {
    const request = new GetCurrentUserRequest();
    return bindNodeCallback<GetCurrentUserResponse>(
      this.client.getCurrentUser.bind(this.client, request, UsersService.authHeader()),
    )();
  }

  oAuthURL(): Observable<string> {
    return this.getCurrentUser().pipe(
      take(1),
      map((resp => {
        const url = new URL(resp.getGoogleOauthConfig().getLoginurl());
        url.searchParams.set('response_type', 'id_token');
        url.searchParams.set('redirect_uri', window.location.origin);
        return url.toString();
      })),
      shareReplay(),
    );
  }
}

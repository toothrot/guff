import {Injectable} from '@angular/core';
import {GetCurrentUserRequest, GetCurrentUserResponse} from '../generated/users_pb';
import {bindNodeCallback, Observable} from 'rxjs';
import {UsersServiceClient} from '../generated/UsersServiceClientPb';
import {HttpClient} from '@angular/common/http';
import {map, shareReplay, take} from 'rxjs/operators';
import {GetDivisionsRequest, GetDivisionsResponse} from '../generated/divisions_pb';
import {DivisionsServiceClient} from '../generated/DivisionsServiceClientPb';

@Injectable({
  providedIn: 'root',
})
export class AppService {
  client: UsersServiceClient;
  divisionsClient: DivisionsServiceClient;

  constructor(private httpClient: HttpClient) {
    this.client = new UsersServiceClient('http://localhost:8080', {}, {});
    this.divisionsClient = new DivisionsServiceClient('http://localhost:8080', {}, {});
  }

  getCurrentUser(): Observable<GetCurrentUserResponse> {
    const request = new GetCurrentUserRequest();
    return bindNodeCallback<GetCurrentUserResponse>(
      this.client.getCurrentUser.bind(this.client, request, {Authorization: 'Bearer ' + localStorage.getItem('token')}),
    )();
  }

  getDivisions(): Observable<GetDivisionsResponse> {
    const request = new GetDivisionsRequest();
    return bindNodeCallback<GetDivisionsResponse>(
      this.divisionsClient.getDivisions.bind(this.divisionsClient, request, {}),
    )();
  }

  oAuthURL(): Observable<string> {
    return this.getCurrentUser().pipe(
      take(1),
      map((resp => {
        const url = new URL(resp.getGoogleOauthConfig().getLoginurl());
        url.searchParams.set('response_type', 'id_token');
        url.searchParams.set('redirect_uri', 'http://localhost:8080');
        return url.toString();
      })),
      shareReplay(),
    );
  }
}

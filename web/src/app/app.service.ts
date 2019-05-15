import {Injectable} from '@angular/core';
import {GetCurrentUserRequest, GetCurrentUserResponse} from '../generated/users_pb';
import {bindNodeCallback, Observable} from 'rxjs';
import {UsersServiceClient} from '../generated/UsersServiceClientPb';
import {HttpClient} from '@angular/common/http';
import {map, shareReplay, take} from 'rxjs/operators';
import {GetDivisionsRequest, GetDivisionsResponse} from '../generated/divisions_pb';
import {DivisionsServiceClient} from '../generated/DivisionsServiceClientPb';
import {ScrapeRequest, ScrapeResponse} from '../generated/admin_pb';
import {AdminServiceClient} from '../generated/AdminServiceClientPb';

@Injectable({
  providedIn: 'root',
})
export class AppService {
  client: UsersServiceClient;
  divisionsClient: DivisionsServiceClient;
  adminClient: AdminServiceClient;

  constructor(private httpClient: HttpClient) {
    this.client = new UsersServiceClient('', {}, {});
    this.divisionsClient = new DivisionsServiceClient('', {}, {});
    this.adminClient = new AdminServiceClient('', {}, {});
  }

  private static authHeader() {
    return {Authorization: 'Bearer ' + localStorage.getItem('token')};
  }

  getCurrentUser(): Observable<GetCurrentUserResponse> {
    const request = new GetCurrentUserRequest();
    return bindNodeCallback<GetCurrentUserResponse>(
      this.client.getCurrentUser.bind(this.client, request, AppService.authHeader()),
    )();
  }

  getDivisions(): Observable<GetDivisionsResponse> {
    const request = new GetDivisionsRequest();
    return bindNodeCallback<GetDivisionsResponse>(
      this.divisionsClient.getDivisions.bind(this.divisionsClient, request, {}),
    )();
  }

  scrapeDivisions(): Observable<ScrapeResponse> {
    const request = new ScrapeRequest();
    return bindNodeCallback<ScrapeResponse>(
      this.adminClient.scrape.bind(this.adminClient, request, AppService.authHeader()),
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

import {Injectable} from '@angular/core';
import {GetCurrentUserRequest, GetCurrentUserResponse} from '../generated/users_pb';
import {bindNodeCallback, Observable} from 'rxjs';
import {UsersServiceClient} from '../generated/UsersServiceClientPb';
import {HttpClient, HttpHeaders, HttpResponse} from '@angular/common/http';
import {first} from 'rxjs/internal/operators/first';
import {tap} from 'rxjs/internal/operators/tap';

@Injectable({
  providedIn: 'root'
})
export class AppService {
  client: UsersServiceClient;

  constructor(private httpClient: HttpClient) {
    this.client = new UsersServiceClient('http://localhost:8080', {}, {});
  }

  getCurrentUser(): Observable<GetCurrentUserResponse> {
    const request = new GetCurrentUserRequest();
    const clientObservable = bindNodeCallback<GetCurrentUserResponse>(
      this.client.getCurrentUser.bind(this.client, request, null)
    );
    return clientObservable();
  }

  logIn(): Observable<HttpResponse<string>> {
    const headers = new HttpHeaders({'Content-Type':  'application/json'});
    return this.httpClient.post('/login', {}, {observe: 'response', responseType: 'text', headers: headers}).pipe(
      first(),
      tap(((resp) => {
        console.log(resp.headers.keys());
      })));
  }
}

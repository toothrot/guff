import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {UsersServiceClient} from 'src/generated/UsersServiceClientPb';
import {Observable, of, throwError} from 'rxjs';
import {mergeMap, tap} from 'rxjs/operators';
import {bindNodeCallback} from 'rxjs/internal/observable/bindNodeCallback';
import {GetCurrentUserRequest, GetCurrentUserResponse} from '../../generated/users_pb';

@Injectable({
  providedIn: 'root',
})
export class OAuthService {
  client: UsersServiceClient;

  constructor(private httpClient: HttpClient) {
    this.client = new UsersServiceClient('http://localhost:8080', {}, {});
  }

  getCurrentUser(idToken = ''): Observable<GetCurrentUserResponse> {
    const request = new GetCurrentUserRequest();
    const clientObservable = bindNodeCallback<GetCurrentUserResponse>(
      this.client.getCurrentUser.bind(this.client, request, {Authorization: 'Bearer ' + idToken}),
    );
    return clientObservable();
  }

  verifyAndSave(idToken = ''): Observable<GetCurrentUserResponse> {
    return this.getCurrentUser(idToken).pipe(
      mergeMap(user => {
        if (!user.getEmail()) {
          localStorage.removeItem('token');
          return throwError(new Error('login failed'));
        }
        return of(user);
      }),
      tap(user => {
        localStorage.setItem('token', idToken);
      }),
    );
  }


}

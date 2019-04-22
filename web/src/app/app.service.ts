import {Injectable} from '@angular/core';
import {GetCurrentUserRequest, GetCurrentUserResponse} from '../generated/users_pb';
import {bindNodeCallback, Observable} from 'rxjs';
import {UsersServiceClient} from '../generated/UsersServiceClientPb';

@Injectable({
  providedIn: 'root'
})
export class AppService {
  client: UsersServiceClient;

  constructor() {
    this.client = new UsersServiceClient('http://localhost:8080', {}, {});
  }

  getCurrentUser(): Observable<GetCurrentUserResponse> {
    const request = new GetCurrentUserRequest();
    const clientObservable = bindNodeCallback<GetCurrentUserResponse>(
      this.client.getCurrentUser.bind(this.client, request, null)
    );
    return clientObservable();
  }
}

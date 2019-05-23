import { Injectable } from '@angular/core';
import {bindNodeCallback, Observable} from 'rxjs';
import {GetDivisionsRequest, GetDivisionsResponse} from '../../generated/divisions_pb';
import {GetTeamsRequest, GetTeamsResponse} from '../../generated/teams_pb';
import {DivisionsServiceClient} from '../../generated/DivisionsServiceClientPb';
import {TeamsServiceClient} from '../../generated/TeamsServiceClientPb';

@Injectable({
  providedIn: 'root'
})
export class DivisionsService {
  private divisionsClient: DivisionsServiceClient;
  private teamsClient: TeamsServiceClient;

  constructor() {
    this.divisionsClient = new DivisionsServiceClient('', {}, {});
    this.teamsClient = new TeamsServiceClient('', {}, {});
  }

  getDivisions(): Observable<GetDivisionsResponse> {
    const request = new GetDivisionsRequest();
    return bindNodeCallback<GetDivisionsResponse>(
      this.divisionsClient.getDivisions.bind(this.divisionsClient, request, {}),
    )();
  }

  getTeams(divisionID = ''): Observable<GetTeamsResponse> {
    const request = new GetTeamsRequest();
    request.setDivisionId(divisionID);
    return bindNodeCallback<GetTeamsResponse>(
      this.teamsClient.getTeams.bind(this.teamsClient, request, {}),
    )();
  }

}

import { Component, OnInit } from '@angular/core';
import {DivisionsService} from './divisions.service';
import {GetDivisionsResponse} from '../../generated/divisions_pb';
import {Observable} from 'rxjs';
import {GetTeamsResponse} from '../../generated/teams_pb';

@Component({
  selector: 'app-division-wall',
  templateUrl: './division-wall.component.html',
  styleUrls: ['./division-wall.component.less']
})
export class DivisionWallComponent implements OnInit {
  private divisions: Observable<GetDivisionsResponse>;
  private teams: Observable<GetTeamsResponse>;

  constructor(private divisionsService: DivisionsService) { }

  ngOnInit() {
    this.divisions = this.divisionsService.getDivisions();
    this.teams = this.divisionsService.getTeams();
  }

}

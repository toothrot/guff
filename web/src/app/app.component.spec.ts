import {async, TestBed} from '@angular/core/testing';
import {RouterTestingModule} from '@angular/router/testing';
import {AppComponent} from './app.component';
import {AppService} from './app.service';
import {OAuthService} from './oauth/oauth.service';
import {Observable, of} from 'rxjs';
import {GetDivisionsResponse} from '../generated/divisions_pb';
import {GetCurrentUserResponse} from '../generated/users_pb';

describe('AppComponent', () => {
  let appServiceStub: Partial<AppService>;
  let oauthServiceStub: Partial<OAuthService>;

  beforeEach(async(() => {
    appServiceStub = {
      oAuthURL(): Observable<string> {
        return of('');
      },
      getDivisions(): Observable<GetDivisionsResponse> {
        return of(new GetDivisionsResponse());
      },
      getCurrentUser(): Observable<GetCurrentUserResponse> {
        return of(new GetCurrentUserResponse());
      },
    };
    oauthServiceStub = {};

    TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
      ],
      declarations: [
        AppComponent,
      ],
      providers: [
        {provide: AppService, useValue: appServiceStub},
        {provide: OAuthService, useValue: oauthServiceStub},
      ],
    }).compileComponents();
  }));

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have as title 'web'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('web');
  });

  it('should render title in a h1 tag', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain('Welcome to web!');
  });
});

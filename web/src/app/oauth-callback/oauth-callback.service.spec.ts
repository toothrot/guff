import {TestBed} from '@angular/core/testing';

import {OAuthCallbackService} from './oauth-callback.service';

describe('OauthCallbackService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: OAuthCallbackService = TestBed.get(OAuthCallbackService);
    expect(service).toBeTruthy();
  });
});

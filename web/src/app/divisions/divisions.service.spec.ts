import { TestBed } from '@angular/core/testing';

import { DivisionsService } from './divisions.service';

describe('DivisionsService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: DivisionsService = TestBed.get(DivisionsService);
    expect(service).toBeTruthy();
  });
});

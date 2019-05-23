import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DivisionWallComponent } from './division-wall.component';

describe('DivisionWallComponent', () => {
  let component: DivisionWallComponent;
  let fixture: ComponentFixture<DivisionWallComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DivisionWallComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DivisionWallComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

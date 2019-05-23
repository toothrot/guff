import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DivisionWallComponent } from './division-wall.component';

@NgModule({
  declarations: [DivisionWallComponent],
  imports: [
    CommonModule
  ],
  exports: [DivisionWallComponent]
})
export class DivisionsModule { }

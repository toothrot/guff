import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {HttpClientModule} from '@angular/common/http';
import {OAuthCallbackComponent} from './oauth-callback/oauth-callback.component';
import {RouterModule, Routes} from '@angular/router';

const routes: Routes = [
  {
    path: 'oauth2callback',
    component: OAuthCallbackComponent,
  },
];

@NgModule({
  declarations: [
    AppComponent,
    OAuthCallbackComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    RouterModule.forRoot(routes),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {
}

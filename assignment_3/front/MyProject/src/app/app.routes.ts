import { Routes } from '@angular/router';
import {AppComponent} from './app.component';
import {ItemComponent} from './components/item/item.component';
import {ItemDetailComponent} from './components/item-detail/item-detail.component';
import {AddItemComponent} from './components/add-item/add-item.component';
import {UpdateItemComponent} from './components/update-item/update-item.component';

export const routes: Routes = [
    {path: '', component: ItemComponent},
    {path: 'item/:id', component: ItemDetailComponent},
    {path: 'add-item', component: AddItemComponent},
    {path: 'update-item/:id', component: UpdateItemComponent},
];

import { Component } from '@angular/core';
import {Item} from '../../models/item';
import {ItemService} from '../../services/item.service';
import {Router} from '@angular/router';
import {FormsModule} from '@angular/forms';

@Component({
  selector: 'app-add-item',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './add-item.component.html',
  styleUrl: './add-item.component.css'
})
export class AddItemComponent {
    item: Item = {id: 0, name: "", price: undefined, rating: undefined, quantity: undefined}
    // item: Item | undefined

    constructor(private itemService: ItemService, private router: Router) {
    }

    addItem() {
        if(this.item) {
            this.itemService.addItem(this.item).subscribe(_ => {
                this.router.navigate(['/']);
            });
        }
    }
}

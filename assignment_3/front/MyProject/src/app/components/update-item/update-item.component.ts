import { Component } from '@angular/core';
import {FormsModule} from '@angular/forms';
import {Item} from '../../models/item';
import {ItemService} from '../../services/item.service';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
  selector: 'app-update-item',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './update-item.component.html',
  styleUrl: './update-item.component.css'
})
export class UpdateItemComponent {
    item: Item = {id: 0, name: "", price: 0, rating: 0.0, quantity: 0}
    // item: Item | undefined

    constructor(private itemService: ItemService, private route: ActivatedRoute, private router: Router) {
    }

    ngOnInit() {
        this.route.paramMap.subscribe(paramMap => {
            const id = paramMap.get('id');

            this.itemService.getItemById(id).subscribe(item => {
                this.item = item;
            }, err => {
                console.log(err);
                this.router.navigate(['/']);
            });
        });
    }

    updateItem() {
        if(this.item) {
            this.itemService.updateItem(this.item.id, this.item).subscribe(item => {
                this.router.navigate(['/item', this.item.id]);
            }, err => {
                console.log(err)
                this.router.navigate(['/item', this.item.id]);
            });
        }
    }
}

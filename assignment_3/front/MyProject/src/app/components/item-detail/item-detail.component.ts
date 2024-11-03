import { Component } from '@angular/core';
import {Item} from '../../models/item';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {ItemService} from '../../services/item.service';

@Component({
  selector: 'app-item-detail',
  standalone: true,
  imports: [
    RouterLink
  ],
  templateUrl: './item-detail.component.html',
  styleUrl: './item-detail.component.css'
})
export class ItemDetailComponent {
    item: Item | undefined;

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
            })
        });
    }

    removeItem(){
        this.itemService.deleteItem(this.item?.id).subscribe(message => {
            console.log(message);
            this.router.navigate(['/']);
        }, err => {
            console.log(err);
        });
    }
}

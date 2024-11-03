import { Component } from '@angular/core';
import {Item} from '../../models/item';
import {ItemService} from '../../services/item.service';
import {NgForOf} from '@angular/common';
import {RouterLink} from '@angular/router';

@Component({
  selector: 'app-item',
  standalone: true,
  imports: [
    NgForOf,
    RouterLink
  ],
  templateUrl: './item.component.html',
  styleUrl: './item.component.css'
})
export class ItemComponent {
    items: Item[] = [];

    constructor(private itemService: ItemService) {
    }

    ngOnInit() {
        this.itemService.getAllItems().subscribe(items => {
            this.items = items;
        }, error => {
            console.log(error);
        });
    }

}

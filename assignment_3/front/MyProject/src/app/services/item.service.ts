import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {Item} from '../models/item';

@Injectable({
  providedIn: 'root'
})
export class ItemService {

    private readonly base_url: string;

    constructor(private http: HttpClient) {
      this.base_url = "http://localhost:8000/items";
    }

    public getAllItems(): Observable<Item[]> {
      return this.http.get<Item[]>(`${this.base_url}`);
    }

    public getItemById(id: any): Observable<Item | any> {
      return this.http.get<Item | any>(`${this.base_url}/${id}`);
    }

    public addItem(item: Item): Observable<Item | any> {
      return this.http.post<Item | any>(`${this.base_url}`, item);
    }

    public updateItem(id: any, item: Item): Observable<Item | any> {
      return this.http.put<Item | any>(`${this.base_url}/${id}`, item);
    }

    public deleteItem(id: any): Observable<any> {
      return this.http.delete<any>(`${this.base_url}/${id}`);
    }
}

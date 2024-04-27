import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProductService } from '../services/product.service';
import { Product } from '../models/product.model';

@Component({
  selector: 'app-product-detail',
  templateUrl: './product-detail.component.html',
  styleUrls: ['./product-detail.component.css']
})
export class ProductDetailComponent implements OnInit {
  product: Product | undefined;
  productId!: number;

  constructor(private route: ActivatedRoute, private productService: ProductService) {}
  
  ngOnInit(): void {
    this.productId = this.route.snapshot.params['id']; // Проверьте, какой параметр используется в вашем маршруте
    this.getProductDetails(this.productId);
  }

  getProductDetails(productId: number): void {
    this.productService.getProductById(productId).subscribe(
      data => {
        this.product = data;
      },
      error => {
        console.error('Error loading product details:', error);
      }
    );
  }
}
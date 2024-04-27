import { Component, OnInit } from '@angular/core';
import { ProductService } from '../services/product.service';
import { Product } from '../models/product.model';
import { Router } from '@angular/router';



@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.css']
})
export class ProductListComponent implements OnInit {
  
  products: Product[]=[];
  

  constructor(private productService: ProductService,
    private router: Router
    
  ) { }


  ngOnInit(): void {
    this.loadProducts();
    
    
  }

  loadProducts(): void {
    this.productService.getAllProducts().subscribe(
      (data: Product[]) => { 
        this.products = data;
        console.log('Products loaded', this.products);
      },
      error => {
        console.error('Error loading products:', error);
      }
    );
  }

  deleteProduct(productId: number): void {
    this.productService.deleteProduct(productId).subscribe(
      () => {
        this.loadProducts();
        console.log('Product deleted');
      },
      error => {
        console.error('Error deleting product:', error);
      }
    );
  }

  navigateToProductDetails(productId: number): void {
    this.router.navigate(['/products', productId]);
  }
  /*
  filterProducts(searchTerm: string): void {
    if (this.productContainers && this.productTitles) {
      this.productContainers.forEach((container: Element) => {
        const productNameElement = container.querySelector(".product-title");
        if (productNameElement) {
          const productName = productNameElement.textContent?.toLowerCase() || '';
          const isVisible = productName.includes(searchTerm.toLowerCase());
          
          if (isVisible) {
            this.renderer.setStyle(container, 'display', 'block');
          } else {
            this.renderer.setStyle(container, 'display', 'none');
          }
        }
      });
    }
*/


}

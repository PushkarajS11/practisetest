<template>
    <div>
            <h1>Prod List</h1>

            <ul>
              <li v-for="(product,x) in products" :key="x">{{product.title}} - {{product.price}}</li>


            </ul>
            
    </div>



</template>


<script>
   
   import shop from '@/api/shop.js'

   export default{
      name : "ProductList",
      data(){

        return {

           products:[]



        }

      } ,

      created() {

          shop.getProducts(products=> {

             this.products=products


          })
         

      }


   }
   

</script>


<style>



</style>


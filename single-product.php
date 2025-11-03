<?php
defined( 'ABSPATH' ) || exit;

remove_action( 'woocommerce_single_product_summary', 'woocommerce_template_single_meta', 40 );
remove_action( 'woocommerce_single_product_summary', 'woocommerce_template_single_price', 10 );

get_header();

if ( ! function_exists( 'wc_get_product' ) ) {
    echo '<p>WooCommerce is not active.</p>';
    get_footer();
    exit;
}
?>

<?php if ( have_posts() ) : while ( have_posts() ) : the_post(); global $product; ?>

  <!-- 🧭 Breadcrumb + Category Row -->

  <div class="product-top-bar">
    <div class="product-top-bar-inner">
      <div class="breadcrumbs">
        <?php
        if ( function_exists( 'woocommerce_breadcrumb' ) ) {
            woocommerce_breadcrumb( array(
                'delimiter'   => ' &raquo; ',
                'wrap_before' => '<nav class="woocommerce-breadcrumb">',
                'wrap_after'  => '</nav>',
            ) );
        }
        ?>
      </div>

      <div class="product-category">
        <div class="product-category-inner">
        <h5>Category </h5>
        <?php
        $terms = wp_get_post_terms( get_the_ID(), 'product_cat' );
        if ( ! empty( $terms ) && ! is_wp_error( $terms ) ) {
            // You can output multiple categories if you wish
            $cats = array();
            foreach ( $terms as $term ) {
                $cats[] = sprintf(
                    '<a href="%s">%s</a>',
                    esc_url( get_term_link( $term ) ),
                    esc_html( $term->name )
                );
            }
            echo implode( ', ', $cats );
        }
        ?>
      </div>
      </div>
    </div>
  </div>

  <!-- 🧾 Product Layout -->
  <div class="single-product-wrapper">
    <div class="single-product-card">

      <div class="single-product-image">
        <?php
        echo $product->get_image( 'large' );

        $gallery_ids = $product->get_gallery_image_ids();
        if ( ! empty( $gallery_ids ) ) {
            echo '<div class="single-product-gallery">';
            foreach ( $gallery_ids as $image_id ) {
                echo wp_get_attachment_image( $image_id, 'thumbnail' );
            }
            echo '</div>';
        }
        ?>
      </div>

      <h1 class="single-product-title"><?php the_title(); ?></h1>

      <p class="single-product-price"><?php echo $product->get_price_html(); ?></p>

      <div class="single-product-description">
        <?php the_content(); ?>
      </div>



    </div>
  </div>

<?php endwhile; endif; ?>

<?php get_footer(); ?>

<?php get_header() ?>


<?php
$params = array(
    'post_type' => 'product',
    'posts_per_page' => -1, // Get all products
);
$wc_query = new WP_Query($params);

if ($wc_query->have_posts()) :
    while ($wc_query->have_posts()) : $wc_query->the_post();
        $product = wc_get_product(get_the_ID()); // Get WooCommerce product object
        echo the_title() . ' - ' . wc_price($product->get_price()); // Display title and price
    endwhile;
    wp_reset_postdata(); // Restore original post data
endif;

?>

<?php get_footer() ?>
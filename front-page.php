<?php get_header(); ?>
<h2 class="front-header">front page</h2>

<div class="product-display-grid">
<?php
$params = array(
    'post_type' => 'product',
    'posts_per_page' => -1,
);
$wc_query = new WP_Query($params);

if ($wc_query->have_posts()) :
    while ($wc_query->have_posts()) : $wc_query->the_post();
        $product = wc_get_product(get_the_ID());
        ?>
        <div class="product-card">
            <a href="<?php the_permalink(); ?>">
                <?php if (has_post_thumbnail()) : ?>
                    <div class="product-image"><?php the_post_thumbnail('medium'); ?></div>
                <?php endif; ?>
                <h3 class="product-title"><?php the_title(); ?></h3>
                <p class="product-price"><?php echo wc_price($product->get_price()); ?></p>
            </a>
        </div>
        <?php
    endwhile;
    wp_reset_postdata();
endif;
?>
</div>

<?php get_footer(); ?>
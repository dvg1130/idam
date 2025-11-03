<?php



function idamiana_enqueue_assets() {
  wp_enqueue_style('idamiana-style', get_stylesheet_uri());
}
add_action('wp_enqueue_scripts', 'idamiana_enqueue_assets');

function tabtitle(){
    add_theme_support('title-tag');
    add_theme_support('post-thumbnails');
}

add_action('after_setup_theme', 'tabtitle');
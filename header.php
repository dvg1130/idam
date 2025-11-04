<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <?php wp_head()?>
</head>
<body>
     <div class="header-wrapper">
        <!-- header logo -->
        <div class="header-left">
            <?php the_custom_logo()?>
        </div>

        <!-- header menu  -->
        <div class="header-right">

            <div class="header-icons">
                <h4>Account</h4>
            </div>

             <div class="header-icons">
                <a href="<?php echo wc_get_cart_url(); ?>" class="cart-icon">
                <!-- SVG icon -->
                <svg class="cart-icon-svg"  xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                     fill="none" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="10" cy="20.5" r="1"/><circle cx="18" cy="20.5" r="1"/>
                    <path d="M2.5 2.5h3l2.7 12.4a2 2 0 0 0 2 1.6h7.7a2 2 0 0 0 2-1.6l1.6-8.4H7.1"/>
                </svg>

                <!-- Item count -->
                <span class="cart-count">
                    <?php echo WC()->cart->get_cart_contents_count(); ?>
                </span>
                </a>

            </div>

             <div class="header-icons">
                <?php
                $attachment_id = 28; // Replace with the actual ID of your image attachment
                $size = array(25,25); // Or 'thumbnail', 'large', 'full', or a custom size array (e.g., array(300, 200))
                $icon = false; // Set to true to display a media icon for non-image attachments
                $attr = array( 'class' => 'my-custom-image-class', 'alt' => 'Descriptive alt text' ); // Optional attributes

                echo wp_get_attachment_image( $attachment_id, $size, $icon, $attr );
                ?>
            </div>


        </div>




     </div>



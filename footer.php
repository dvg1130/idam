<div class="footer-wrapper">


    <div class="footer-content-left">
        <div class="social-list">
            <div class="social-logo">
                <?php
                $attachment_id = 26; // Replace with the actual ID of your image attachment
                $size = array(25,25); // Or 'thumbnail', 'large', 'full', or a custom size array (e.g., array(300, 200))
                $icon = false; // Set to true to display a media icon for non-image attachments
                $attr = array( 'class' => 'my-custom-image-class', 'alt' => 'Descriptive alt text' ); // Optional attributes

                echo wp_get_attachment_image( $attachment_id, $size, $icon, $attr );
                ?>
            </div>

            <div class="social-logo">
                <?php
                $attachment_id = 25; // Replace with the actual ID of your image attachment
                $size = array(25,25); // Or 'thumbnail', 'large', 'full', or a custom size array (e.g., array(300, 200))
                $icon = false; // Set to true to display a media icon for non-image attachments
                $attr = array( 'class' => 'my-custom-image-class', 'alt' => 'Descriptive alt text' ); // Optional attributes

                echo wp_get_attachment_image( $attachment_id, $size, $icon, $attr );
                ?>
            </div>

            <div class="social-logo">
                <?php
                $attachment_id = 24; // Replace with the actual ID of your image attachment
                $size = array(25,25); // Or 'thumbnail', 'large', 'full', or a custom size array (e.g., array(300, 200))
                $icon = false; // Set to true to display a media icon for non-image attachments
                $attr = array( 'class' => 'my-custom-image-class', 'alt' => 'Descriptive alt text' ); // Optional attributes

                echo wp_get_attachment_image( $attachment_id, $size, $icon, $attr );
                ?>
            </div>


        </div>
    </div>

     <div class="footer-content-center">
        <div class="footer-logo">
        <?php the_custom_logo()?>
    </div>
    </div>

    <div class="footer-content-right">
    <h3>link</h3>
    <h3>link</h3>
    </div>

</div>
<?php wp_footer() ?>
</body>
</html>
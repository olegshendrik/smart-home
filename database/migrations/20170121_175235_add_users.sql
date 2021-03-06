-- +migrate Up

INSERT INTO `roles`(`name`,`description`,`parent`,`created_at`,`update_at`) VALUES ( 'demo', '', NULL, '2017-01-15 05:17:09', '2017-01-15 05:17:09' );
INSERT INTO `roles`(`name`,`description`,`parent`,`created_at`,`update_at`) VALUES ( 'admin', '', NULL, '2017-01-15 05:20:58', '2017-01-15 05:21:29' );

INSERT INTO `users`(`id`,`nickname`,`first_name`,`last_name`,`encrypted_password`,`email`,`history`,`status`,`reset_password_token`,`authentication_token`,`image_id`,`sign_in_count`,`current_sign_in_ip`,`last_sign_in_ip`,`user_id`,`role_name`,`reset_password_sent_at`,`current_sign_in_at`,`last_sign_in_at`,`created_at`,`update_at`,`deleted`) VALUES ( '1', 'admin', '', '', 'f6fdffe48c908deb0f4c3bd36c032e72', 'admin@e154.ru', '[]', 'active', '', 'xlzEaHNBbn80OmTfWd1z18XpNUlZikdb4z5fo5YAxlNv3CfWxs', NULL, '111', '127.0.0.1', '127.0.0.1', NULL, 'admin', NULL, '2017-01-21 11:11:26', '2017-01-21 10:54:03', '2017-01-15 05:25:07', '2017-01-21 11:11:26', NULL );
INSERT INTO `users`(`id`,`nickname`,`first_name`,`last_name`,`encrypted_password`,`email`,`history`,`status`,`reset_password_token`,`authentication_token`,`image_id`,`sign_in_count`,`current_sign_in_ip`,`last_sign_in_ip`,`user_id`,`role_name`,`reset_password_sent_at`,`current_sign_in_at`,`last_sign_in_at`,`created_at`,`update_at`,`deleted`) VALUES ( '2', 'demo', '', '', 'c514c91e4ed341f263e458d44b3bb0a7', 'demo@e154.ru', '[]', 'active', '', '5SLTHOzN1hWw6jhgEw0y9JbtwdBIK5mgW3DLt5FYy23zNkVnvW', NULL, '8', '127.0.0.1', '127.0.0.1', NULL, 'demo', NULL, '2017-01-21 11:11:43', '2017-01-20 17:28:23', '2017-01-18 17:13:28', '2017-01-21 11:11:43', NULL );

INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '49', '1', 'phone1', '' );
INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '50', '1', 'phone2', '' );
INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '51', '1', 'phone3', '' );
INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '52', '1', 'autograph', '' );
INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '53', '2', 'phone1', '' );
INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '54', '2', 'phone2', '' );
INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '55', '2', 'phone3', '' );
INSERT INTO `user_metas`(`id`,`user_id`,`key`,`value`) VALUES ( '56', '2', 'autograph', '' );

INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '1', 'demo', 'dashboard', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '2', 'demo', 'dashboard', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '3', 'demo', 'dashboard', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '4', 'demo', 'dashboard', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '5', 'demo', 'flow', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '6', 'demo', 'flow', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '7', 'demo', 'flow', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '8', 'demo', 'flow', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '9', 'demo', 'device', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '10', 'demo', 'device', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '11', 'demo', 'device', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '12', 'demo', 'device', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '13', 'demo', 'scenarios', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '14', 'demo', 'scenarios', 'create_script' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '15', 'demo', 'scenarios', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '16', 'demo', 'scenarios', 'delete_script' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '17', 'demo', 'scenarios', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '18', 'demo', 'scenarios', 'read_script' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '19', 'demo', 'scenarios', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '20', 'demo', 'scenarios', 'update_script' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '21', 'demo', 'ws', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '22', 'demo', 'workflow', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '23', 'demo', 'workflow', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '24', 'demo', 'workflow', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '25', 'demo', 'workflow', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '26', 'demo', 'device_action', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '27', 'demo', 'device_action', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '28', 'demo', 'device_action', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '29', 'demo', 'device_action', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '30', 'demo', 'device_state', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '31', 'demo', 'device_state', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '32', 'demo', 'device_state', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '33', 'demo', 'device_state', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '34', 'demo', 'log', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '35', 'demo', 'notifr', 'create_notifr_template' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '36', 'demo', 'notifr', 'read_notifr_template' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '37', 'demo', 'notifr', 'repeat_notify' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '38', 'demo', 'notifr', 'show_notify' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '39', 'demo', 'notifr', 'update_notifr_template' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '40', 'demo', 'notifr', 'create_notifr_item' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '41', 'demo', 'notifr', 'delete_notifr_item' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '42', 'demo', 'notifr', 'delete_notifr_template' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '43', 'demo', 'notifr', 'delete_notify' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '44', 'demo', 'notifr', 'preview_notifr' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '45', 'demo', 'notifr', 'read_notifr_item' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '46', 'demo', 'notifr', 'update_notifr_item' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '47', 'demo', 'notifr', 'create_notify' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '48', 'demo', 'user', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '49', 'demo', 'user', 'read_role' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '50', 'demo', 'user', 'read_role_access_list' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '51', 'demo', 'worker', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '52', 'demo', 'worker', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '53', 'demo', 'worker', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '54', 'demo', 'worker', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '55', 'demo', 'image', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '56', 'demo', 'image', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '57', 'demo', 'image', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '58', 'demo', 'image', 'upload' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '59', 'demo', 'image', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '60', 'demo', 'map', 'delete_map' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '61', 'demo', 'map', 'delete_map_layer' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '62', 'demo', 'map', 'read_map' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '63', 'demo', 'map', 'read_map_element' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '64', 'demo', 'map', 'update_map_layer' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '65', 'demo', 'map', 'create_map_element' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '66', 'demo', 'map', 'create_map_layer' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '67', 'demo', 'map', 'delete_map_element' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '68', 'demo', 'map', 'read_map_layer' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '69', 'demo', 'map', 'update_map' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '70', 'demo', 'map', 'update_map_element' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '71', 'demo', 'map', 'create_map' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '72', 'demo', 'node', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '73', 'demo', 'node', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '74', 'demo', 'node', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '75', 'demo', 'node', 'update' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '76', 'demo', 'script', 'create' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '77', 'demo', 'script', 'delete' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '78', 'demo', 'script', 'exec_script' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '79', 'demo', 'script', 'read' );
INSERT INTO `permissions`(`id`,`role_name`,`package_name`,`level_name`) VALUES ( '80', 'demo', 'script', 'update' );

-- +migrate Down
DELETE FROM `user_metas` WHere user_id IN (1,2);
DELETE FROM `users` WHERE  id in (1,2);
DELETE FROM permissions WHERE role_name IN ("admin", "user", "demo");
DELETE FROM roles WHERE  name in ("admin");
DELETE FROM roles WHERE  name in ("user");
DELETE FROM roles WHERE  name in ("demo");


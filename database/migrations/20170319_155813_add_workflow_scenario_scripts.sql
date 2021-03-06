-- +migrate Up
CREATE TABLE workflow_scenario_scripts (
id Int( 22 ) AUTO_INCREMENT NOT NULL,
workflow_scenario_id Int( 22 ) NOT NULL,
script_id Int( 22 ) NOT NULL,
PRIMARY KEY ( id ),
CONSTRAINT unique_id UNIQUE( id ) )
ENGINE = InnoDB;

ALTER TABLE `workflow_scenario_scripts` ADD CONSTRAINT `lnk_scripts_workflow_scenario_scripts` FOREIGN KEY ( `script_id` ) REFERENCES `scripts`( `id` ) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `workflow_scenario_scripts` ADD CONSTRAINT `lnk_workflow_scenarios_workflow_scenario_scripts` FOREIGN KEY (`workflow_scenario_id`) REFERENCES `workflow_scenarios` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE `workflow_scenario_scripts` DROP FOREIGN KEY `lnk_scripts_workflow_scenario_scripts`;
ALTER TABLE `workflow_scenario_scripts` DROP FOREIGN KEY `lnk_workflow_scenarios_workflow_scenario_scripts`;
DROP TABLE IF EXISTS `workflow_scenario_scripts` CASCADE;


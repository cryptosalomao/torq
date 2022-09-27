import React from 'react';

import {
  ChevronDown20Regular as CollapsedIcon,
  LineHorizontal120Regular as ExpandedIcon,
  FluentIconsProps,
} from '@fluentui/react-icons';
import classNames from 'classnames';

import styles from './sectionContainer.module.scss';

type SectionContainerProps = {
  title: string;
  icon: React.FC<FluentIconsProps>;
  children: React.ReactNode;
  expanded?: boolean;
  disabled?: boolean;
  handleToggle?: (event: React.MouseEvent<HTMLDivElement, MouseEvent>) => void;
};

export function SectionContainer(props: SectionContainerProps) {
  return (
    <div className={classNames(styles.sectionContainer, { [styles.disabled]: props.disabled })}>
      <div className={styles.sectionTitleContainer} onClick={props.handleToggle}>
        <div className={styles.sidebarIcon}>
          <props.icon />
        </div>
        <div className={styles.sidebarTitle}>{props.title}</div>
        <div className={styles.sidebarIcon}>{props.expanded ? <ExpandedIcon /> : <CollapsedIcon />}</div>
      </div>
      <div className={classNames(styles.sidebarSectionContent, { [styles.expanded]: props.expanded })}>
        {props.children}
      </div>
    </div>
  );
}

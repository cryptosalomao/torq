import React from 'react';

import classNames from 'classnames';

import Breadcrumbs from 'features/breadcrumbs/Breadcrumbs';

import styles from './templates.module.scss';


type PageTitleProps = {
  title: string;
  breadcrumbs?: Array<any>;
  className?: string;
  children?: React.ReactNode;
};

function PageTitle(props: PageTitleProps) {
  return (
    <div className={classNames(styles.pageTitleWrapper, props.className)}>
      <div className={styles.leftWrapper}>
        <Breadcrumbs breadcrumbs={props.breadcrumbs || []} />
        <h1 className={styles.titleContainer}>{props.title}</h1>
      </div>
      {props.children && <div className={styles.rightWrapper}>{props.children}</div>}
    </div>
  );
}

const memoizedPageTitle = React.memo(PageTitle);
export default memoizedPageTitle;

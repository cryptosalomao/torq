import React from 'react';

import classNames from 'classnames';

import PageTitle from 'features/templates/PageTitle';

import styles from './details-page-template.module.scss';


type DetailsPageProps = {
  title: string;
  titleContent?: React.ReactNode;
  sidebarExpanded?: boolean;
  sidebar?: React.ReactNode;
  breadcrumbs?: Array<any>;
  children?: React.ReactNode;
};

function DetailsPage(props: DetailsPageProps) {
  return (
    <div className={styles.contentWrapper}>
      <PageTitle breadcrumbs={props.breadcrumbs} title={props.title} className={styles.detailsPageTitle}>
        {props.titleContent}
      </PageTitle>

      <div className={styles.detailsPageContent}>{props.children}</div>

      <div className={classNames(styles.pageSidebarWrapper, { [styles.sidebarExpanded]: props.sidebarExpanded })}>
        {props.sidebar}
      </div>
    </div>
  );
}

const memoizedDetailsPage = React.memo(DetailsPage);
export default memoizedDetailsPage;

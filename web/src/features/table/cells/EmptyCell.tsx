import classNames from 'classnames';

import styles from './cell.module.scss';

function EmptyCell(index?: number | string, className?: string) {
  return <div className={classNames(styles.cell, styles.empty, className)} key={'last-cell-' + index} />;
}

export default EmptyCell;

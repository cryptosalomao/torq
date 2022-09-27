import { useState } from 'react';

import { MoneyHand24Regular as TransactionIconModal } from '@fluentui/react-icons';

import ChannelPage from 'features/channel/ChannelPage';
import PopoutPageTemplate from 'features/templates/popoutPageTemplate/PopoutPageTemplate';

type PopoutChannelsPageProps = {
  show: boolean;
  modalCloseHandler: () => void;
};

function PopoutChannelsPage(props: PopoutChannelsPageProps) {
  const [expandAdvancedOptions, setExpandAdvancedOptions] = useState(false);

  // const handleAdvancedToggle = () => {
  //   setExpandAdvancedOptions(!expandAdvancedOptions);
  // };

  return (
    <PopoutPageTemplate
      title={'Channel'}
      show={props.show}
      onClose={props.modalCloseHandler}
      icon={<TransactionIconModal />}
    >
      <ChannelPage />
    </PopoutPageTemplate>
  );
}

export default PopoutChannelsPage;

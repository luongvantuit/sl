if [[ ! -f "$PWD/install-fabric.sh" ]]; then
    curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x $PWD/install-fabric.sh

    $PWD/install-fabric.sh
fi

find $PWD/fabric-samples/bin -maxdepth 0 -empty -exec $PWD/install-fabric.sh \;

export PATH=$PWD/fabric-samples/bin:$PATH
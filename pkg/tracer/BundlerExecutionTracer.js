var tracer = {
  reverts: [],
  validationOOG: false,
  executionOOG: false,
  executionGasLimit: 0,

  _depth: 0,
  _executionGasStack: [],
  _defaultGasItem: { used: 0, required: 0 },
  _marker: 0,
  _validationMarker: 1,
  _executionMarker: 3,
  _aiOperationEventTopics0:
    "0x7b592a4f684d032578f75dff2ca3d3d2aec981a6d0a782e1d127595a6511a4f1",

  _isValidation: function () {
    return (
      this._marker >= this._validationMarker &&
      this._marker < this._executionMarker
    );
  },

  _isExecution: function () {
    return this._marker === this._executionMarker;
  },

  _isAiOperationEvent: function (log) {
    var topics0 = "0x" + log.stack.peek(2).toString(16);
    return topics0 === this._aiOperationEventTopics0;
  },

  _setAiOperationEvent: function (opcode, log) {
    var count = parseInt(opcode.substring(3));
    var ofs = parseInt(log.stack.peek(0).toString());
    var len = parseInt(log.stack.peek(1).toString());
    var topics = [];
    for (var i = 0; i < count; i++) {
      topics.push("0x" + log.stack.peek(2 + i).toString(16));
    }
    var data = toHex(log.memory.slice(ofs, ofs + len));
    this.aiOperationEvent = {
      topics: topics,
      data: data,
    };
  },

  fault: function fault(log, db) {},
  result: function result(ctx, db) {
    return {
      reverts: this.reverts,
      validationOOG: this.validationOOG,
      executionOOG: this.executionOOG,
      executionGasLimit: this.executionGasLimit,
      aiOperationEvent: this.aiOperationEvent,
      output: toHex(ctx.output),
      error: ctx.error,
    };
  },

  enter: function enter(frame) {
    if (this._isExecution()) {
      var next = this._depth + 1;
      if (this._executionGasStack[next] === undefined)
        this._executionGasStack[next] = Object.assign({}, this._defaultGasItem);
    }
  },
  exit: function exit(frame) {
    if (this._isExecution()) {
      if (frame.getError() !== undefined) {
        this.reverts.push(toHex(frame.getOutput()));
      }

      if (this._depth >= 2) {
        // Get the final gas item for the nested frame.
        var nested = Object.assign(
          {},
          this._executionGasStack[this._depth + 1] || this._defaultGasItem
        );

        // Reset the nested gas item to prevent double counting on re-entry.
        this._executionGasStack[this._depth + 1] = Object.assign(
          {},
          this._defaultGasItem
        );

        // Keep track of the total gas used by all frames at this depth.
        // This does not account for the gas required due to the 63/64 rule.
        var used = frame.getGasUsed();
        this._executionGasStack[this._depth].used += used;

        // Keep track of the total gas required by all frames at this depth.
        // This accounts for additional gas needed due to the 63/64 rule.
        this._executionGasStack[this._depth].required +=
          used - nested.used + Math.ceil((nested.required * 64) / 63);

        // Keep track of the final gas limit.
        this.executionGasLimit = this._executionGasStack[this._depth].required;
      }
    }
  },

  step: function step(log, db) {
    var opcode = log.op.toString();
    this._depth = log.getDepth();
    if (this._depth === 1 && opcode === "NUMBER") this._marker++;

    if (
      this._depth <= 2 &&
      opcode.startsWith("LOG") &&
      this._isAiOperationEvent(log)
    )
      this._setAiOperationEvent(opcode, log);

    if (log.getGas() < log.getCost() && this._isValidation())
      this.validationOOG = true;

    if (log.getGas() < log.getCost() && this._isExecution())
      this.executionOOG = true;
  },
};
